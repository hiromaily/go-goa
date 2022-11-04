import React from 'react'
import Button from '@mui/material/Button'

type ModeButtonProps = {
  isDarkMode: boolean
  onClick: () => Promise<void>
}

const buttonName = new Map<boolean, string>([
  [false, 'Dark Mode'],
  [true, 'Light Mode'],
])

const ModeButton = ({ isDarkMode, onClick }: ModeButtonProps) => {
  return (
    <>
      <Button
        color='inherit'
        variant='outlined'
        size='medium'
        sx={{ minWidth: 120 }}
        onClick={onClick}
      >
        {buttonName.get(isDarkMode)}
      </Button>
    </>
  )
}

export default ModeButton
