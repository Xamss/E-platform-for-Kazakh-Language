import React from 'react';
import Typography from '@mui/material/Typography';
import Box from '@mui/material/Box';
import Grid from '@mui/material/Grid';
import Stack from '@mui/material/Stack';
import HeroImage from './Images/yurt.png';

const Hero = () => {
  return (
    <section>
      {/* mobile screen */}
      <Stack sx={{ display: { xs: 'flex', md: 'none' }, px: 2 }}>
        <Box>
          <img
            src={HeroImage}
            alt="presentation"
            style={{
              maxWidth: '100%',
              height: 'auto',
            }}
          />
        </Box>
        <Stack alignItems="center" sx={{ my: 3 }}>
          <Box sx={{ mb: 3, textAlign: 'center' }}>
            <Typography
              variant="h3"
              component="h1"
              sx={{
                color: 'neutral.veryDarkViolet',
                mb: 1.5,
                letterSpacing: -1,
                lineHeight: 1,
                fontWeight: 800,
                fontSize: '2.35rem',
              }}
            >
              The Kazakh Steppe
            </Typography>
            <Typography
              variant="body2"
              component="p"
              sx={{
                color: 'neutral.grayishViolet',
                lineHeight: 1.73,
                fontSize: '0.92rem',
              }}
            >
              Dive into interactive language lessons that bring the Kazakh language to life. With easy-to-follow grammar tutorials, fun vocabulary drills, and clear pronunciation guides, learning Kazakh has never been more accessible or enjoyable.
              Join us today and start your adventure into the heart of Kazakhstan!
            </Typography>
          </Box>
        
        </Stack>
      </Stack>

      {/* desktop screen */}
      <Grid
        container
        direction="row"
        sx={{
          display: { xs: 'none', md: 'flex' },
          px: 6,
          py: 8,
          mx: 'auto',
          overflow: 'hidden',
        }}
      >
        <Grid item md={7}>
          <Box sx={{ my: 5, maxWidth: '85%' }}>
            <Typography
              variant="h1"
              component="h1"
              sx={{
                color: 'neutral.veryDarkViolet',
                fontSize: '5rem',
                mb: 1.5,
                letterSpacing: -3,
                lineHeight: 0.9,
              }}
            >
              The Kazakh Steppe
            </Typography>
            <Typography
              variant="body1"
              component="p"
              sx={{
                color: 'neutral.grayishViolet',
              }}
            >
              Dive into interactive language lessons that bring the Kazakh language to life. With easy-to-follow grammar tutorials, fun vocabulary drills, and clear pronunciation guides, learning Kazakh has never been more accessible or enjoyable.
              Join us today and start your adventure into the heart of Kazakhstan!
            </Typography>
          </Box>
          
        </Grid>

        <Grid item md={5}>
          <img
            src={HeroImage}
            alt="presentation"
            style={{
              maxWidth: '130%',
              height: 'auto',
            }}
          />
        </Grid>
      </Grid>
    </section>
  );
};

export default Hero;
