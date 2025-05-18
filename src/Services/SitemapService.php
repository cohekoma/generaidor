<?php

namespace Cohekoma\Generaidor\Services;

use Cohekoma\Generaidor\Core\Paths as Paths;

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
        $path = Paths::CONTENT_DIR->getFullPath();

        // Scan the directory where all posts are stored.
        $contentDirScanned = scandir($path, SCANDIR_SORT_NONE);
        $contentDir = array_diff($contentDirScanned, array('..', '.'));

//        $di = new \RecursiveDirectoryIterator($path);
//        foreach (new \RecursiveIteratorIterator($di) as $filename => $file) {
//            echo $filename . ' - ' . $file->getSize() . ' bytes ' . "\n";
//        }

        $iterator = new \RecursiveIteratorIterator(
            new \RecursiveDirectoryIterator($path, \RecursiveDirectoryIterator::SKIP_DOTS),
            \RecursiveIteratorIterator::SELF_FIRST
        );

        foreach ($iterator as $fileInfo) {
            var_dump($fileInfo);
        }

        return $result;
    }
}