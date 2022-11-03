#!/bin/bash

# eslint, pretter
npm install --save-dev eslint-plugin-import @typescript-eslint/parser eslint-import-resolver-typescript
npm install --save-dev prettier eslint-config-prettier eslint-plugin-prettier eslint-plugin-react-hooks
touch .prettierrc
# {
#   "semi": false,
#   "tabWidth": 2,
#   "printWidth": 100,
#   "singleQuote": true,
#   "trailingComma": "all",
#   "jsxSingleQuote": true,
#   "bracketSpacing": true
# }

# directory
mkdir -p src/components
mv pages src/
mv styles src/

# files
rm -rf src/pages/api
rm -rf src/styles/Home.module.css

mv src/pages/index.tsx src/pages/index.default.tsx
touch src/pages/index.tsx

touch src/styles/index.module.scss

# semantic UI React
# https://react.semantic-ui.com/usage/
#npm install semantic-ui-react semantic-ui-css

# mui
npm install @mui/material @emotion/react @emotion/styled
npm install @mui/icons-material


# others
npm install sass
npm install styled-components
npm install --save-dev @types/webpack-env
npm install --save-dev @next/bundle-analyzer