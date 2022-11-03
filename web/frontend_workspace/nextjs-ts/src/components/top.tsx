import Box from '@mui/material/Box'
import Container from '@mui/material/Container'
import { useWindowSize } from 'usehooks-ts'

type TopProps = {
  message: string
}

const Top = ({ message }: TopProps) => {
  const { width, height } = useWindowSize()

  return (
    <Container
      sx={{
        minHeight: height > 135 ? height - 135 + 'px' : '100px',
        height: 'auto !important',
      }}
    >
      <p>{message}</p>
    </Container>
  )
}

export default Top
