<?php

declare(strict_types=1);

namespace Cohekoma\Generaidor\Services;

/**
 * Core service that is used to build the whole static site.
 */
class SiteBuilderService {

    /**
     * Generate Sitemap, then use the Sitemap to extract content & metadata in each Markdown file.
     * For each file, map the returned content with template, then finally output an HTML file.
     * @return void
     */
    public function build() : void
    {
        // Generated Sitemap. Prefer to use SitemapService
        $pages = [];

        // Loop through all pages retrieved from sitemap.
        foreach ($pages as $page) {
            // Parse markdown content. May return an associated arr with two items: yaml/metedata & actual post content
            // Map the returned content with the layout in template.
            // Output / write the final content as HTML file.
        }

        echo "Star building the site...";
    }
}
