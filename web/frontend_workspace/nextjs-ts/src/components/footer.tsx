import * as React from 'react'
import FavoriteIcon from '@mui/icons-material/Favorite'
import LocationOnIcon from '@mui/icons-material/LocationOn'
import RestoreIcon from '@mui/icons-material/Restore'
import BottomNavigation from '@mui/material/BottomNavigation'
import BottomNavigationAction from '@mui/material/BottomNavigationAction'
import Box from '@mui/material/Box'
import { useDarkMode } from 'usehooks-ts'

const Footer = () => {
  const [value, setValue] = React.useState(0)
  const { isDarkMode } = useDarkMode()
  const iconColor = isDarkMode ? 'secondary' : 'inherit'

  console.log(`iconColor: ${iconColor}`)
  return (
    <Box border={1}>
      <BottomNavigation
        sx={{ color: 'inherit' }}
        showLabels
        value={value}
        onChange={(event, newValue) => {
          setValue(newValue)
        }}
      >
        <BottomNavigationAction
          sx={{ color: iconColor }}
          label='Recents'
          icon={<RestoreIcon />}
        />
        <BottomNavigationAction
          sx={{ color: iconColor }}
          label='Favorites'
          icon={<FavoriteIcon />}
        />
        <BottomNavigationAction
          sx={{ color: iconColor }}
          label='Nearby'
          icon={<LocationOnIcon />}
        />
      </BottomNavigation>
    </Box>
  )
}

export default Footer
