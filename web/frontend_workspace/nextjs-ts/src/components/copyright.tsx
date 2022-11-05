import * as React from 'react'
import Link from '@mui/material/Link'
import Typography from '@mui/material/Typography'

const Copyright = (props: any) => {
  return (
    <Typography variant='body2' color='text.secondary' align='center' {...props}>
      {'Copyright © '}
      <Link color='inherit' href='https://github.com/hiromaily/go-goa'>
        go-goa
      </Link>{' '}
      {new Date().getFullYear()}
      {'.'}
    </Typography>
  )
}

export { Copyright }
