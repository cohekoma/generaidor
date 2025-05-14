#!/usr/bin/env php

<?php
require_once __DIR__ . '/vendor/autoload.php';

use Cohekoma\Generaidor\Parser\Yaml\YamlParser as YamlParser;

$sample = new YamlParser();
var_dump($sample->parse());
