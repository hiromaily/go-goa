import { createTheme } from '@mui/material/styles'

export const monotoneTheme = createTheme({
  palette: {
    mode: 'light',
    primary: {
      main: 'rgba(0,0,0,0.57)',
      light: 'rgba(0,0,0,0.33)',
      dark: 'rgba(0,0,0,0.83)',
    },
    secondary: {
      main: 'rgba(8,8,8,0.69)',
      light: 'rgba(57,57,57,0.36)',
    },
  },
})
