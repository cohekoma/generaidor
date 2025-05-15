<?php

namespace Cohekoma\Generaidor\Core;

enum Paths
{
    case CONTENT_DIR;
    case DIST_DIR;
    case TEMPLATE_DIR;

    public function getFullPath(): string
    {
        $rootDir = dirname(__FILE__, 3);
        return match($this) {
            self::CONTENT_DIR => $rootDir . "/content",
            self::DIST_DIR => $rootDir . "/dist",
            self::TEMPLATE_DIR => $rootDir . "/templates"

        };
    }
}