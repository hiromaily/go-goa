import '../styles/globals.css'
import type { AppProps } from 'next/app'
import CssBaseline from '@mui/material/CssBaseline'
import { ThemeProvider } from '@mui/material/styles'
import { useDarkMode } from 'usehooks-ts'
import { monotoneTheme } from '../theme/theme'

if (process.env.NODE_ENV === 'development' || process.env.NODE_ENV === 'test') {
  const mockServer = () => import('@/mocks/worker')
  mockServer()
}

function MyApp({ Component, pageProps }: AppProps) {
  const { isDarkMode } = useDarkMode()

  // switch theme by isDarkMode Flag
  let theme = monotoneTheme
  if (isDarkMode) theme.palette.mode = 'dark'
  else theme.palette.mode = 'light'

  return (
    <ThemeProvider theme={theme}>
      <CssBaseline />
      <Component {...pageProps} />
    </ThemeProvider>
  )
}

export default MyApp
