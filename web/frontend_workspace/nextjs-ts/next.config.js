const BundleAnalyzer = require("@next/bundle-analyzer");

// https://github.com/vercel/next.js/tree/canary/packages/next-bundle-analyzer
const withBundleAnalyzer = BundleAnalyzer({
  enabled: process.env.ANALYZE === "true",
});

/** @type {import('next').NextConfig} */
const nextConfig = {
  reactStrictMode: true,
  swcMinify: true,
}

module.exports = withBundleAnalyzer(nextConfig)
