<?php

namespace Cohekoma\Generaidor\Core;

enum Paths : string
{
    case CONTENT_DIR = "content";
    case DIST_DIR = "dist";
    case TEMPLATE_DIR = "templates";

    public function getFullPath(): string
    {
        return dirname(__FILE__, 3) . "/" . $this->value;
    }
}