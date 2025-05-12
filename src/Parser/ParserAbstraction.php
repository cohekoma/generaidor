<?php

namespace Cohekoma\Generaidor\Parser;

abstract class ParserAbstraction implements ParserInterface
{
    public function parse() : void
    {
        echo "Parsing?";
    }
}