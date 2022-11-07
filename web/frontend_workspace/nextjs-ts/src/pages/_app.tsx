import '../styles/globals.css'
import type { AppProps } from 'next/app'
import type { EmotionCache } from '@emotion/cache'
import { CacheProvider } from '@emotion/react'
import CssBaseline from '@mui/material/CssBaseline'
import { ThemeProvider } from '@mui/material/styles'
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
  let dynamicDarkMode = isDarkMode

  // Set 'dark' as default mode for server side because current my environment is dark mode in client side
  // This setting can fix `Prop 'className' did not match` Error
  if (typeof window === 'undefined') {
    // server side
    dynamicDarkMode = true
  }

  // switch theme by isDarkMode Flag
  const theme = monotoneTheme(dynamicDarkMode)

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
