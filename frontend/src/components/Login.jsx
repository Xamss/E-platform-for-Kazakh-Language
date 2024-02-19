import Button from 'react-bootstrap/Button';
import Card from 'react-bootstrap/Card';
import Form from 'react-bootstrap/Form';
import Container from 'react-bootstrap/Container';
import React, { useEffect, useState, useContext } from 'react';
import { toast } from 'react-toastify';
import axios from 'axios';
import { useNavigate } from 'react-router-dom';
import { FaUserAlt, FaKey, FaSignInAlt } from 'react-icons/fa';
import UserContext from '../context/UserContext';
import '../App.css';
import './Footer.css';

function Login() {
  const [username, setUsername] = useState('');
  const [password, setPassword] = useState('');
  const URL = `${import.meta.env.VITE_API_ENDPOINT}`;
  const navigate = useNavigate();
  const context = useContext(UserContext);
  let toastId = null;

  useEffect(() => {
    const token = localStorage.getItem('token');
    if (token) {
      toast.warning('You are already logged in');
      setTimeout(() => {
        navigate('/');
      }, 4000);
    }
  }, [navigate]);

  const handleSubmit = async (e) => {
    e.preventDefault();
    toastId = null;
    toastId = toast.loading('Logging in...');
    if (!validateForm()) return;

    await fetchLogin();
  };

  function validateForm() {
    if (username === '' || password === '') {
      toast.update(toastId, {
        render: 'Please fill all the fields',
        type: 'error',
        isLoading: false,
        autoClose: 2000,
      });
      return false;
    }
    return true;
  }

  const fetchLogin = async () => {
    try {
      console.log(URL)
      const response = await axios.post(`${URL}/user/login`, {
        username,
        password,
      });
      const {
        token,
      } = response.data;

      //    set token and logged-in user's name in local storage
      localStorage.setItem('token', token);
      localStorage.setItem('name', username);
      context.setUser({ token: token });
      toast.update(toastId, {
        render: `Welcome ${username}`,
        type: 'success',
        isLoading: false,
        autoClose: 2000,
      });
      navigate('/');
    } catch (
      {
      response: {
        status,
        data: { error },
      },
    }
    ) {
      toast.update(toastId, {
        render: 'Login failed',
        type: 'error',
        isLoading: false,
        autoClose: 2000,
      });
      if (status === 400) {
        toast.error(error);
      } else if (status === 401) {
        toast.error(error);
      } else {
        toast.error('Something went wrong');
      }
    }
  };

  return (
    <Container
      className="d-flex align-items-center justify-content-center"
      style={{ minHeight: '90vh' }}
    >
      <div className="w-100" style={{ maxWidth: '400px' }}>
        <Card>
          <Card.Header style={{ backgroundColor: '#538d45' }}>
            {' '}
            <h4 style={{ backgroundColor: '#538d45' }}>Login</h4>{' '}
          </Card.Header>
          <Card.Body>
            <Form onSubmit={handleSubmit}>
              <Form.Group className="mb-3" controlId="formBasicEmail">
                <Form.Label
                  style={{ display: 'flex', gap: '5px', alignItems: 'center' }}
                >
                  <FaUserAlt /> Username
                </Form.Label>
                <Form.Control
                  placeholder="Enter username"
                  onChange={(e) => setUsername(e.target.value)}
                  aria-required={true}
                />
                <Form.Text className="text-muted">
                  We'll never share your email with anyone else.
                </Form.Text>
              </Form.Group>
              <Form.Group className="mb-3" controlId="formBasicPassword">
                <Form.Label
                  style={{ display: 'flex', gap: '5px', alignItems: 'center' }}
                >
                  <FaKey /> Password
                </Form.Label>
                <Form.Control
                  type="password"
                  placeholder="Password"
                  onChange={(e) => setPassword(e.target.value)}
                  aria-required={true}
                />
                <a
                  onClick={() => {
                    navigate('/forgot-password');
                  }}
                  style={{
                    textDecoration: 'none',
                    color: '#538d45',
                    cursor: 'pointer',
                  }}
                >
                    Forgot password
                </a>{' '}
              </Form.Group>
              <Button
                className={'w-100'}
                variant="info"
                type="submit"
                style={{ backgroundColor: '#538d45', color: 'white' }}
              >
                <FaSignInAlt style={{ marginRight: '0.3rem' }} />
                LOGIN
              </Button>
            </Form>
          </Card.Body>
          <Card.Footer
            className="text-muted"
            style={{ display: 'flex', justifyContent: 'space-between' }}
          >
            Don't Have an Account?{' '}
            <a
              style={{
                textDecoration: 'none',
                color: '#538d45',
                cursor: 'pointer',
              }}
              onClick={() => {
                navigate('/signup');
              }}
            >
              Click Here to Signup
            </a>
          </Card.Footer>
        </Card>
      </div>
    </Container>
  );
}

export default Login;
