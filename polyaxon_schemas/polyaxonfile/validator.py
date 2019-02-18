# -*- coding: utf-8 -*-
from __future__ import absolute_import, division, print_function

import copy

from polyaxon_schemas.build import BuildConfig
from polyaxon_schemas.exceptions import PolyaxonfileError
from polyaxon_schemas.hptuning import HPTuningConfig
from polyaxon_schemas.logging import LoggingConfig
from polyaxon_schemas.ml.eval import EvalConfig
from polyaxon_schemas.ml.models import ModelConfig
from polyaxon_schemas.ml.train import TrainConfig
from polyaxon_schemas.run_exec import RunConfig


def validate_headers(spec, data):
    """Validates headers data and creates the config objects"""
    validated_data = {
        spec.VERSION: data[spec.VERSION],
        spec.KIND: data[spec.KIND],
    }

    if data.get(spec.LOGGING):
        validated_data[spec.LOGGING] = LoggingConfig.from_dict(
            data[spec.LOGGING])

    if data.get(spec.TAGS):
        validated_data[spec.TAGS] = data[spec.TAGS]

    if data.get(spec.HP_TUNING):
        validated_data[spec.HP_TUNING] = HPTuningConfig.from_dict(
            data[spec.HP_TUNING])

    return validated_data


def validate(spec, data):
    """Validates the data and creates the config objects"""
    data = copy.deepcopy(data)
    validated_data = {}

    def validate_keys(section, config, section_data):
        if not isinstance(section_data, dict) or section == spec.MODEL:
            return

        extra_args = [key for key in section_data.keys() if key not in config.SCHEMA().fields]
        if extra_args:
            raise PolyaxonfileError('Extra arguments passed for `{}`: {}'.format(
                section, extra_args))

    def add_validated_section(section, config):
        if data.get(section):
            section_data = data[section]
            validate_keys(section=section, config=config, section_data=section_data)
            validated_data[section] = config.from_dict(section_data)

    add_validated_section(spec.ENVIRONMENT, spec.ENVIRONMENT_CONFIG)
    add_validated_section(spec.BUILD, BuildConfig)
    add_validated_section(spec.RUN, RunConfig)
    add_validated_section(spec.MODEL, ModelConfig)
    add_validated_section(spec.TRAIN, TrainConfig)
    add_validated_section(spec.EVAL, EvalConfig)

    return validated_data
