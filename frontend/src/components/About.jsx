import Accordion from 'react-bootstrap/Accordion';

import React from 'react';

function About() {
  return (
    <div className="d-flex justify-content-center">
      <div className="w-75 p-3" style={{ maxWidth: '200px;' }}>
        <h1>About YerBakh</h1>
        <p>
        There's more to explore beyond the language! Our website takes you on a fascinating tour of Kazakhstan's history, 
        from its ancient nomadic roots to its vibrant modern identity. Discover the beauty of traditional Kazakh music, 
        literature, and art. Get a taste of the nation's exquisite cuisine and learn about its unique customs and celebrations.
        </p>
        <h4>Why Use YerBakh?</h4>
        <p>
          YerBakh is an open-source language learning platform with quizzes, cultural videos, 
          and forums for sharing your learning journey with fellow enthusiasts, 
          our platform is more than just a learning tool—it's a gateway to understanding 
          and appreciating the rich tapestry of Kazakh life.
        </p>
        <h4>Features</h4>
        <p>
          YerBakh provides a variety of features to improve your web experience.
          These include:
        </p>
        <ol>
          <li>Interactive Map of Kazakhstan</li>
          <li>Wide-range of exercises</li>
          <li>Progress Tracking</li>
          <li>AI-powered Learning Experience</li>
          <li>Quick Introduction to Kazakh Culture</li>
        </ol>
        <br />
        <h3>Frequently Asked Questions</h3>
        <br />
        <Accordion className="px-md-3">
          <Accordion.Item eventKey="0">
            <Accordion.Header>1. Question example</Accordion.Header>
            <Accordion.Body>
            Lorem ipsum dolor sit amet, consectetur adipiscing elit. Phasellus a arcu et ligula viverra porttitor. Proin a ipsum vitae tellus mollis vestibulum at non enim. In volutpat tincidunt elit, vitae accumsan elit hendrerit in. Cras vitae fermentum lacus, hendrerit eleifend lacus. Maecenas metus nisi, mollis vel laoreet in, semper a diam. Donec accumsan augue mi, id tincidunt urna imperdiet et. Nullam mollis sem ac eleifend ullamcorper.
            </Accordion.Body>
          </Accordion.Item>
          <Accordion.Item eventKey="1">
            <Accordion.Header>
              2. Question example
            </Accordion.Header>
            <Accordion.Body>
            Lorem ipsum dolor sit amet, consectetur adipiscing elit. Phasellus a arcu et ligula viverra porttitor. Proin a ipsum vitae tellus mollis vestibulum at non enim. In volutpat tincidunt elit, vitae accumsan elit hendrerit in. Cras vitae fermentum lacus, hendrerit eleifend lacus. Maecenas metus nisi, mollis vel laoreet in, semper a diam. Donec accumsan augue mi, id tincidunt urna imperdiet et. Nullam mollis sem ac eleifend ullamcorper.
            </Accordion.Body>
          </Accordion.Item>
          <Accordion.Item eventKey="2">
            <Accordion.Header>3. Question example</Accordion.Header>
            <Accordion.Body>
            Lorem ipsum dolor sit amet, consectetur adipiscing elit. Phasellus a arcu et ligula viverra porttitor. Proin a ipsum vitae tellus mollis vestibulum at non enim. In volutpat tincidunt elit, vitae accumsan elit hendrerit in. Cras vitae fermentum lacus, hendrerit eleifend lacus. Maecenas metus nisi, mollis vel laoreet in, semper a diam. Donec accumsan augue mi, id tincidunt urna imperdiet et. Nullam mollis sem ac eleifend ullamcorper.
            </Accordion.Body>
          </Accordion.Item>
          <Accordion.Item eventKey="3">
            <Accordion.Header>
              4. Question example
            </Accordion.Header>
            <Accordion.Body>
            Lorem ipsum dolor sit amet, consectetur adipiscing elit. Phasellus a arcu et ligula viverra porttitor. Proin a ipsum vitae tellus mollis vestibulum at non enim. In volutpat tincidunt elit, vitae accumsan elit hendrerit in. Cras vitae fermentum lacus, hendrerit eleifend lacus. Maecenas metus nisi, mollis vel laoreet in, semper a diam. Donec accumsan augue mi, id tincidunt urna imperdiet et. Nullam mollis sem ac eleifend ullamcorper.
            </Accordion.Body>
          </Accordion.Item>
        </Accordion>
      </div>
    </div>
  );
}

export default About;
