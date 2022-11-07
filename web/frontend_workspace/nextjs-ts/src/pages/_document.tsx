import Document, {
  DocumentContext,
  DocumentProps,
  Html,
  Head,
  Main,
  NextScript,
} from 'next/document'
import * as React from 'react'
import { EmotionJSX } from '@emotion/react/types/jsx-namespace'
import createEmotionServer from '@emotion/server/create-instance'
import createEmotionCache from '../utils/createEmotionCache'

// Refer to
// - https://github.com/search?q=getStaticProps+createEmotionServer&type=code
// - https://github.com/aumapichai/Storymap/blob/eb0ae7f4677a4daa022513cf0387e5abbddd33cd/src/pages/_document.tsx
// - https://github.com/VityaSchel/the-archive/blob/4bbb07748275ba60ffb446c7be1250bcad5b528f/pages/_document.tsx
// - https://github.com/oubakiou/100perts/blob/58bf550eb5778079ca2d7b809f0cd51a64cb65d1/v0.1.1/helloworld-app/docs/chapter3.md
//   - Material-UI を導入してみようの項

type MyDocumentProps = DocumentProps & {
  emotionStyleTags: EmotionJSX.Element[]
}

const MyDocument = (props: MyDocumentProps) => (
  <Html>
    <Head>
      <link
        rel='stylesheet'
        href='https://fonts.googleapis.com/css?family=Roboto:300,400,500,700&display=swap'
      />
      <link rel='stylesheet' href='https://fonts.googleapis.com/icon?family=Material+Icons' />
      <link
        rel='stylesheet'
        href='https://cdn.jsdelivr.net/gh/lipis/flag-icons@6.6.6/css/flag-icons.min.css'
      />
      {props.emotionStyleTags}
    </Head>
    <body>
      <Main />
      <NextScript />
    </body>
  </Html>
)

// https://github.com/mui/material-ui/blob/master/examples/nextjs/pages/_document.js
MyDocument.getStaticProps = async (ctx: DocumentContext) => {
  // eslint-disable-next-line  testing-library/render-result-naming-convention
  const originalRenderPage = ctx.renderPage

  const cache = createEmotionCache()
  const { extractCriticalToChunks } = createEmotionServer(cache)

  /* eslint-disable */
  ctx.renderPage = () =>
    originalRenderPage({
      enhanceApp: (App: any) =>
        function EnhanceApp(props) {
          // FIXME: error by `tsc --noEmit`
          return <App emotionCache={cache} {...props} />
        },
    })
  /* eslint-enable */

  const initialProps = await Document.getInitialProps(ctx)

  const emotionStyles = extractCriticalToChunks(initialProps.html)
  const emotionStyleTags = emotionStyles.styles.map((style) => (
    <style
      data-emotion={`${style.key} ${style.ids.join(' ')}`}
      key={style.key}
      // eslint-disable-next-line react/no-danger
      dangerouslySetInnerHTML={{ __html: style.css }}
    />
  ))

  return {
    ...initialProps,
    emotionStyleTags,
  }
}

export default MyDocument
