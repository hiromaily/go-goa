import dynamic from 'next/dynamic'
import Head from 'next/head'
import Image from 'next/image'
import Router from 'next/router'
import AppBar from '@mui/material/AppBar'
import Box from '@mui/material/Box'
import Button from '@mui/material/Button'
import Toolbar from '@mui/material/Toolbar'
import Typography from '@mui/material/Typography'
import { useDarkMode } from 'usehooks-ts'
import { spacing } from '@mui/system';
//import ModeButton from './buttons/mode'
// FIX: Warning: Text content did not match
const DynamicMuteButton = dynamic(() => import('./buttons/mode'), { ssr: false })

type HeaderProps = {
  title: string
}

const Header = ({ title }: HeaderProps) => {
  //const [isDarkMode, setDarkMode] = useState(false);
  const { isDarkMode, toggle } = useDarkMode()

  const onClickTop = () => {
    Router.push('/')
  }
  const onClickAdmin = () => {
    Router.push('/admin')
  }

  // switch mode by button
  const onClickMode = async () => {
    console.log(`current isDarkMode: ${isDarkMode}`)
    toggle()
  }

  // useEffect(() => {
  //   //mounted
  //   const isDark = useMediaQuery('(prefers-color-scheme: dark)')
  //   setDarkMode(isDark)
  // }, [])

  return (
    <>
      <Head>
        <title>{title}</title>
        <meta name='viewport' content='initial-scale=1.0, width=device-width' />
        <link rel='icon' href='/favicon.ico' />
      </Head>
      <AppBar sx={{ color: 'primary' }} position='static'>
        <Toolbar>
          <Box sx={{ marginRight: '20px' }}>
            <Image src='/hy.png' alt='HY Logo' width={40} height={40} />
          </Box>
          <Typography variant='h6' component='div' sx={{ flexGrow: 1 }}>
            go-goa example site
          </Typography>
          <Button color='inherit' onClick={onClickTop}>
            Top
          </Button>
          <Button color='inherit' onClick={onClickAdmin}>
            Admin
          </Button>
          <Button sx={{ mr: 1 }} color='inherit' variant='outlined'>
            Login
          </Button>
          <DynamicMuteButton isDarkMode={isDarkMode} onClick={onClickMode} />
        </Toolbar>
      </AppBar>
    </>
  )
}

export default Header
