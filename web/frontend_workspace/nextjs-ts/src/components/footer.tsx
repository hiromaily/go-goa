// import Link from 'next/link'
// import Style from '../styles/common.module.scss'

import * as React from 'react'
import FavoriteIcon from '@mui/icons-material/Favorite'
import LocationOnIcon from '@mui/icons-material/LocationOn'
import RestoreIcon from '@mui/icons-material/Restore'
import BottomNavigation from '@mui/material/BottomNavigation'
import BottomNavigationAction from '@mui/material/BottomNavigationAction'
import Box from '@mui/material/Box'

const Footer = () => {
  const [value, setValue] = React.useState(0)

  return (
    <Box>
      <BottomNavigation
        sx={{ bgcolor: '#808080' }}
        showLabels
        value={value}
        onChange={(event, newValue) => {
          setValue(newValue)
        }}
      >
        <BottomNavigationAction sx={{ color: 'white' }} label='Recents' icon={<RestoreIcon />} />
        <BottomNavigationAction sx={{ color: 'white' }} label='Favorites' icon={<FavoriteIcon />} />
        <BottomNavigationAction sx={{ color: 'white' }} label='Nearby' icon={<LocationOnIcon />} />
      </BottomNavigation>
    </Box>
  )
}

export default Footer
