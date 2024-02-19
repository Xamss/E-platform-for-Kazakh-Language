import React from 'react';
import { Container } from '@mui/material';
import Hero from './Hero';


const HomePage = () => {
  return (
    <Container maxWidth="xl" sx={{ pr: 0 }}>
      <Hero />

      <div style={{ margin: '4rem' }}></div>
    </Container>
  );
};

export default HomePage;
