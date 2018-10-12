# -*- coding: utf-8 -*-
from __future__ import absolute_import, division, print_function

import requests
import time

from unittest import TestCase

from polyaxon_client.api_config import ApiConfig
from polyaxon_client.transport.periodic_transport import PeriodicTransportMixin
from polyaxon_client.workers.periodic_worker import PeriodicWorker


class DummyTransport(PeriodicTransportMixin):
    # pylint:disable=protected-access
    def __init__(self, delay=0):
        self.queue = []
        self.delay = delay
        self.config = ApiConfig(in_cluster=True, timeout=0.001, interval=0)
        self._periodic_exceptions = 0
        self._periodic_done = 0

    def post(self, url, **kwargs):
        time.sleep(self.delay)
        self.queue.append(('post', url, kwargs))


class ExceptionTransport(PeriodicTransportMixin):
    # pylint:disable=protected-access
    def __init__(self, delay=0):
        self.delay = delay
        self.config = ApiConfig(in_cluster=True, timeout=0.001, interval=0)
        self._periodic_exceptions = 0
        self._periodic_done = 0

    def post(self, **kwargs):
        time.sleep(self.delay)
        raise requests.exceptions.HTTPError('error')


class TestPeriodicTransport(TestCase):
    # pylint:disable=protected-access
    def setUp(self):
        self.transport = DummyTransport()
        self.exception_transport = ExceptionTransport()

    def test_retry_session(self):
        assert hasattr(self.transport, '_retry_session') is False
        assert isinstance(self.transport.retry_session, requests.Session)
        assert isinstance(self.transport._retry_session, requests.Session)

    def test_worker(self):
        assert hasattr(self.transport, '_periodic_workers') is False
        assert isinstance(self.transport.periodic_workers, dict)
        assert isinstance(self.transport._periodic_workers, dict)
        assert isinstance(self.transport.get_periodic_worker('foo', request=self.transport.post),
                          PeriodicWorker)

    def test_periodic_requests_with_data(self):
        assert self.transport.queue == []

        self.transport.periodic_post(url='url_post', data={'d1': 'v1'})
        time.sleep(0.001)
        queue = self.transport.queue
        assert len(queue) == 1
        assert queue[0][0] == 'post'
        assert queue[0][1] == 'url_post'
        assert queue[0][2]['data'] == [{'d1': 'v1'}]
        assert self.transport.done == 1
        assert self.transport.exceptions == 0
        assert self.transport.get_periodic_worker('url_post').is_alive() is True

        self.transport.periodic_post(url='url_post', data={'d1': 'v1'})
        self.transport.periodic_post(url='url_post', data={'d2': 'v2'})
        time.sleep(0.001)
        queue = self.transport.queue
        assert len(queue) == 2
        assert queue[1][0] == 'post'
        assert queue[1][1] == 'url_post'
        assert queue[1][2]['data'] == [{'d1': 'v1'}, {'d2': 'v2'}]
        assert self.transport.done == 2
        assert self.transport.exceptions == 0
        assert self.transport.get_periodic_worker('url_post').is_alive() is True

    def test_periodic_requests_with_json_data(self):
        assert self.transport.queue == []

        self.transport.periodic_post(url='url_post', json_data={'d1': 'v1'})
        time.sleep(0.001)
        queue = self.transport.queue
        assert len(queue) == 1
        assert queue[0][0] == 'post'
        assert queue[0][1] == 'url_post'
        assert queue[0][2]['json_data'] == [{'d1': 'v1'}]
        assert self.transport.done == 1
        assert self.transport.exceptions == 0
        assert self.transport.get_periodic_worker('url_post').is_alive() is True

        self.transport.periodic_post(url='url_post', json_data={'d1': 'v1'})
        self.transport.periodic_post(url='url_post', json_data={'d2': 'v2'})
        time.sleep(0.001)
        queue = self.transport.queue
        assert len(queue) == 2
        assert queue[1][0] == 'post'
        assert queue[1][1] == 'url_post'
        assert queue[1][2]['json_data'] == [{'d1': 'v1'}, {'d2': 'v2'}]
        assert self.transport.done == 2
        assert self.transport.exceptions == 0
        assert self.transport.get_periodic_worker('url_post').is_alive() is True

    def test_periodic_requests_with_different_urls(self):
        self.transport.periodic_post(url='url_post1', json_data={'d11': 'v11'})
        self.transport.periodic_post(url='url_post1', json_data={'d12': 'v12'})
        self.transport.periodic_post(url='url_post2', data={'d21': 'v21'})
        self.transport.periodic_post(url='url_post2', data={'d22': 'v22'})
        time.sleep(0.001)
        queue = self.transport.queue
        assert len(queue) == 2
        assert queue[0][0] == 'post'
        assert queue[1][0] == 'post'
        assert queue[0][1] == 'url_post1'
        assert queue[1][1] == 'url_post2'
        assert queue[0][2]['json_data'] == [{'d11': 'v11'}, {'d12': 'v12'}]
        assert queue[1][2]['data'] == [{'d21': 'v21'}, {'d22': 'v22'}]
        assert self.transport.done == 2
        assert self.transport.exceptions == 0
        assert self.transport.get_periodic_worker('url_post1').is_alive() is True
        assert self.transport.get_periodic_worker('url_post2').is_alive() is True

    def test_async_exceptions(self):
        self.exception_transport.periodic_post(url='url_post', data={'d1': 'v1'})
        time.sleep(2)
        assert self.exception_transport.done == 1
        assert self.exception_transport.exceptions == 1

    def test_worker_atexit_handle_queue_before_stopping(self):
        # Transport
        self.transport.config.timeout = 0.5
        self.transport.delay = 0.5
        assert self.transport.queue == []
        self.transport.periodic_post(url='url_post', data={'d1': 'v1'})
        assert self.transport.queue == []
        time.sleep(0.001)
        assert self.transport.get_periodic_worker('url_post').is_alive() is True
        self.transport.get_periodic_worker('url_post').atexit()
        queue = self.transport.queue
        assert len(queue) == 1
        assert queue[0][0] == 'post'
        assert queue[0][1] == 'url_post'
        assert queue[0][2]['data'] == [{'d1': 'v1'}]
        assert self.transport.done == 1
        assert self.transport.exceptions == 0
        assert self.transport.periodic_workers.get('url_post').is_alive() is False

        # Exception transport
        self.exception_transport.config.timeout = 0.5
        self.exception_transport.delay = 0.5
        self.exception_transport.periodic_post(url='url_post', data={'d1': 'v1'})
        assert self.exception_transport.done == 0
        assert self.exception_transport.exceptions == 0
        time.sleep(0.1)
        assert self.exception_transport.get_periodic_worker('url_post').is_alive() is True
        self.exception_transport.get_periodic_worker('url_post').atexit()
        assert self.exception_transport.done == 1
        assert self.exception_transport.exceptions == 1
        assert self.exception_transport.periodic_workers.get('url_post').is_alive() is False
