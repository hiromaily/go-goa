import React, { ReactNode } from 'react'
import Card from '@mui/material/Card'

type BaseCardProps = {
  children?: ReactNode
}

const BaseCard = ({ children }: BaseCardProps) => (
  <Card
    sx={{
      marginBottom: '30px',
    }}
  >
    {children}
  </Card>
)

export default BaseCard
