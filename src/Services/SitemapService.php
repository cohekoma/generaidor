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

            /**
             * @TODO: At this part right here, for each file path, we already have a flat array with the file name is the last item in the array. What we need to do here is to think about how to construct a tree-like array from a simple flat array, so that we can add into the final result. We can do like this: 1) Pop the last item. 2) Reverse the array. 3) Now, we will going backwards (from the bottom up), so we can wrap each element around the last item that we just popped.
             */
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