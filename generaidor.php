#!/usr/bin/env php

<?php
require_once __DIR__ . '/vendor/autoload.php';

use Cohekoma\Generaidor\Parser\Yaml\YamlParser as YamlParser;
const  MARKDOWN_CONTENT_DIR = '/content';

$allContentFiles = glob(__DIR__ . MARKDOWN_CONTENT_DIR . '/*.md');
$samplePost = $allContentFiles[1];
$samplePostContent = file_get_contents($samplePost);

$metaDataRegEx = '/^---([\s\S]*?)---/';
preg_match($metaDataRegEx, $samplePostContent, $matches);

$metaData = explode("\n", trim($matches[1]));
var_dump($metaData);

$sample = new YamlParser();
$sample->parse();
