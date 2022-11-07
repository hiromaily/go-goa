import React, { ReactNode } from 'react'
import CardHeader from '@mui/material/CardHeader'
import { boxColor, fontColor } from '../../theme/theme'

type JobHeaderProps = {
  isDarkMode: boolean
  title: string
  flag: string
  //children?: ReactNode
}

const JobHeader = ({ isDarkMode, title, flag }: JobHeaderProps) => {
  const flagData = `fi fi-${flag}`
  return (
    <CardHeader
      sx={{
        height: '45px',
        color: fontColor(isDarkMode),
        backgroundColor: boxColor(isDarkMode),
      }}
      avatar={<span className={flagData}></span>}
      titleTypographyProps={{ variant: 'h6' }}
      title={title}
    />
  )
}

export default JobHeader
