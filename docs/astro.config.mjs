import { defineConfig } from 'astro/config';
import starlight from '@astrojs/starlight';
import starlightLlmsTxt from 'starlight-llms-txt';
import pagefind from 'astro-pagefind';
import swup from '@swup/astro';
import compress from '@playform/compress';
import brokenLinksChecker from 'astro-broken-links-checker';

// https://astro.build/config
export default defineConfig({
	site: 'https://netlify.github.io',
	base: '/gocommerce',
	integrations: [
		starlight({
			title: 'GoCommerce',
			sidebar: [
				{
					label: 'Getting Started',
					items: [
						{ label: 'Introduction', link: '/introduction/' },
						{ label: 'Quickstart', link: '/quickstart/' },
					],
				},
				{
					label: 'Guides',
					items: [
						{ label: 'Deep Dive', link: '/guides/deep-dive/' },
						{ label: 'Multi Instances', link: '/guides/multi-instances/' },
					],
				},
				{
					label: 'API Reference',
					link: '/api-reference/',
				},
			],
			plugins: [
				starlightLlmsTxt(),
			],
		}),
		pagefind(),
		swup(),
		compress(),
        brokenLinksChecker(),
	],
});
