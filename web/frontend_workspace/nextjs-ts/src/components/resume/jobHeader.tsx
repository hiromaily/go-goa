import React, { ReactNode } from 'react'
import CardHeader from '@mui/material/CardHeader'
import { boxColor, fontColor } from '../../theme/theme'

type JobHeaderProps = {
  isDarkMode: boolean
  title: string
  //children?: ReactNode
}

const JobHeader = ({ isDarkMode, title }: JobHeaderProps) => (
  <CardHeader
    sx={{
      height: '45px',
      color: fontColor(isDarkMode),
      backgroundColor: boxColor(isDarkMode),
    }}
    titleTypographyProps={{ variant: 'h6' }}
    title={title}
  />
)

export default JobHeader
