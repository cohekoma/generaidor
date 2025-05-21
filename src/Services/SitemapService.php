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
            // remove the beginning slash so when exploding, the result will be more accurate (var_dump to check).
            $filePathName = explode("/", ltrim(str_replace($contentDirPath, "", $filename), "/"));

            if (sizeof($filePathName) == 1) {
                $result[] = $filePathName[0];
            } else {
                var_dump(sizeof($filePathName));
            }

            var_dump($filePathName);
        }

        var_dump($result);
        return $result;
    }
}