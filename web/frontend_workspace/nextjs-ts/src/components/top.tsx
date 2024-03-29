import Link from 'next/link'
import React from 'react'
import Box from '@mui/material/Box'
import Container from '@mui/material/Container'
import { useWindowSize } from 'usehooks-ts'

type TopProps = {
  message: string
}

// eslint-disable-next-line react/display-name
const Top = React.forwardRef(({ message }: TopProps, _ref) => {
  const { width, height } = useWindowSize()

  return (
    <Container
      sx={{
        minHeight: height > 120 ? height - 120 + 'px' : '100px',
        height: 'auto !important',
      }}
    >
      <Box pt={4}>
        <ul>
          <li>
            <Link href='/'>Home</Link>
          </li>
          <li>
            <Link href='/login'>Login</Link>
          </li>
          <li>
            <Link href='/admin'>Admin</Link>
          </li>
          <li>
            <Link href='/resume'>Resume</Link>
          </li>
        </ul>
        <p>{message}</p>
      </Box>
    </Container>
  )
})

export default Top
