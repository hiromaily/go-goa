import { createTheme } from '@mui/material/styles'

const monotoneTheme = (isDarkMode: boolean) => {
  console.log(`after changed isDarkMode: ${isDarkMode}`)
  const mode = isDarkMode ? 'dark' : 'light'

  return createTheme({
    palette: {
      mode: mode,
      primary: {
        main: 'rgba(0,0,0,0.5)',
        light: 'rgba(0,0,0,0.3)',
        dark: 'rgba(0,0,0,0.8)',
      },
      secondary: {
        main: 'rgba(0,0,0,0.5)',
        light: 'rgba(0,0,0,0.3)',
        dark: 'rgba(0,0,0,0.3)',
      },
    },
  })
}

const boxColor = (isDarkMode: boolean): string => {
  if (isDarkMode) return 'black'
  else return 'darkgray'
}

const fontColor = (isDarkMode: boolean): string => {
  return 'white'
}

export { monotoneTheme, boxColor, fontColor }
