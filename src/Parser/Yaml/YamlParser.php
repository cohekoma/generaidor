<?php

namespace Cohekoma\Generaidor\Parser\Yaml;

use Cohekoma\Generaidor\Parser\ParserAbstraction;
use Cohekoma\Generaidor\Core\Paths as Paths;

class YamlParser extends ParserAbstraction {

    /**
     * @var string $metaDataRegEx Regex pattern to extract meta data from markdown file
     */
    private string $metaDataRegEx = '/^---([\s\S]*?)---/';

    public function parse() : array|string
    {
        $allContentFiles = glob(Paths::CONTENT_DIR->getFullPath() . '/*.md');
        $samplePost = $allContentFiles[1];
        $samplePostContent = file_get_contents($samplePost);

        preg_match($this->metaDataRegEx, $samplePostContent, $matches);

        return explode("\n", trim($matches[1]));
    }
}