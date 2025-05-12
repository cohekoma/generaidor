<?php

namespace generaidor\Parser;
interface ParserInterface
{
    /*
     * Parse the content into acceptable type
     */
    function parse();
}

abstract class ParserAbstraction implements ParserInterface
{
    public function parse()
    {
        echo "Parsing?";
    }
}