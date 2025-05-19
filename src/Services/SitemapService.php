<?php

namespace Cohekoma\Generaidor\Services;

use Cohekoma\Generaidor\Core\Paths as Paths;
use FilesystemIterator;

/**
 * SitemapService is a service that does one main job: return a sitemap.
 */
class SitemapService {

    /**
     * Generate the sitemap based on the Markdown files within the content folder (or any equivalent folder).
     * @return array
     */
    public function generate() : array
    {
        $result = [];
        $contentDirPath = Paths::CONTENT_DIR->getFullPath();

        // Scan the directory where all posts are stored.
        $contentDir = new \RecursiveDirectoryIterator($contentDirPath, FilesystemIterator::SKIP_DOTS);
        foreach (new \RecursiveIteratorIterator($contentDir) as $filename => $file) {
            $filePathName = str_replace(ROOT_PATH, "", $filename);
            var_dump($filePathName);
        }

        return $result;
    }
}