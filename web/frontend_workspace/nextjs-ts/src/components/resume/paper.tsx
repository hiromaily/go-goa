import React, { ReactNode } from 'react'
import Paper from '@mui/material/Paper'
import { paperBgColor } from '../../theme/theme'
//import { styled } from '@mui/material/styles'

type BasePaperProps = {
  isDarkMode: boolean
  children?: ReactNode
}

// const Item = styled(Paper)(({ theme }) => ({
//   backgroundColor: theme.palette.mode === 'dark' ? '#1A2027' : '#fff',
//   ...theme.typography.body2,
//   padding: theme.spacing(1),
//   textAlign: 'center',
//   color: theme.palette.text.secondary,
// }))

const BasePaper = ({ isDarkMode, children }: BasePaperProps) => (
  <Paper
    sx={{
      backgroundColor: paperBgColor(isDarkMode),
      padding: '4px',
      textAlign: 'center',
    }}
  >
    {children}
  </Paper>
)

export default BasePaper
