import '../styles/globals.css'
import type { AppProps } from 'next/app'
import type { EmotionCache } from '@emotion/cache'
import { CacheProvider } from '@emotion/react'
import CssBaseline from '@mui/material/CssBaseline'
import { ThemeProvider } from '@mui/material/styles'
//import PropTypes from 'prop-types'
import { useDarkMode } from 'usehooks-ts'
import { monotoneTheme } from '../theme/theme'
import createEmotionCache from '../utils/createEmotionCache'

if (process.env.NODE_ENV === 'development' || process.env.NODE_ENV === 'test') {
  const mockServer = () => import('@/mocks/worker')
  mockServer()
}

const clientSideEmotionCache = createEmotionCache()

interface MyAppProps extends AppProps {
  emotionCache?: EmotionCache
}

function MyApp({ Component, emotionCache = clientSideEmotionCache, pageProps }: MyAppProps) {
  const { isDarkMode } = useDarkMode()
  // switch theme by isDarkMode Flag
  const theme = monotoneTheme(isDarkMode)

  return (
    <CacheProvider value={emotionCache}>
      <ThemeProvider theme={theme}>
        <CssBaseline />
        <Component {...pageProps} />
      </ThemeProvider>
    </CacheProvider>
  )
}

// MyApp.propTypes = {
//   Component: PropTypes.elementType.isRequired,
//   emotionCache: PropTypes.object,
//   pageProps: PropTypes.object.isRequired,
// };

export default MyApp
