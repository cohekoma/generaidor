#!/usr/bin/env php

<?php

const  MARKDOWN_CONTENT_DIR = '/content';

$allContentFiles = glob(__DIR__ . MARKDOWN_CONTENT_DIR . '/*.md');
$samplePost = $allContentFiles[1];
$samplePostContent = file_get_contents($samplePost);

$metaDataRegEx = '/^---([\s\S]*?)---/';
preg_match($metaDataRegEx, $samplePostContent, $matches);

$metaData = explode("\n", trim($matches[1]));
var_dump($metaData);
