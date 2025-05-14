<?php

namespace Cohekoma\Generaidor\Parser;
interface ParserInterface
{
    /*
     * Parse the content into acceptable type
     */
    function parse() : array|string;
}