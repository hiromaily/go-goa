import Head from 'next/head'
import Image from 'next/image'
import Router from 'next/router'
import AppBar from '@mui/material/AppBar'
import Box from '@mui/material/Box'
import Button from '@mui/material/Button'
import Toolbar from '@mui/material/Toolbar'
import Typography from '@mui/material/Typography'

type HeaderProps = {
  title: string
}

const Header = ({ title }: HeaderProps) => {
  const onClickTop = () => {
    Router.push('/')
  }
  const onClickAdmin = () => {
    Router.push('/admin')
  }

  return (
    <>
      <Head>
        <title>{title}</title>
        <meta name='viewport' content='initial-scale=1.0, width=device-width' />
        <link rel='icon' href='/favicon.ico' />
      </Head>
      <AppBar sx={{ bgcolor: '#808080' }} position='static'>
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
          <Button variant='outlined'>Login</Button>
        </Toolbar>
      </AppBar>
    </>
  )
}

export default Header
