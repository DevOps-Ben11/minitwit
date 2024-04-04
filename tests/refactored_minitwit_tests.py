# -*- coding: utf-8 -*-
"""
    MiniTwit Tests
    ~~~~~~~~~~~~~~

    Tests a MiniTwit application.

    :refactored: (c) 2024 by HelgeCPH from Armin Ronacher's original unittest version
    :copyright: (c) 2010 by Armin Ronacher.
    :license: BSD, see LICENSE for more details.
"""
import requests
import json


# import schema
# import data
# otherwise use the database that you got previously
BASE_URL = "http://localhost:5000/api"

def register(username, password, email=None):
    """Helper function to register a user"""
    if email is None:
        email = username + '@example.com'
    return requests.post(f'{BASE_URL}/register', data=json.dumps({
        'username':     username,
        'password':     password,
        'email':        email,
    }), allow_redirects=True)

def login(username, password):
    """Helper function to login"""
    http_session = requests.Session()
    r = http_session.post(f'{BASE_URL}/login', data=json.dumps({
        'username': username,
        'password': password
    }), allow_redirects=True)
    return r, http_session

def register_and_login(username, password):
    """Registers and logs in in one go"""
    register(username, password)
    return login(username, password)

def logout(http_session):
    """Helper function to logout"""
    return http_session.get(f'{BASE_URL}/logout', allow_redirects=True)

def add_message(http_session, text):
    """Records a message"""
    r = http_session.post(f'{BASE_URL}/add_message',
        data=json.dumps({'message': text}),
        allow_redirects=True
    )
    if text:
        assert r.ok
    return r

def test_register():
    """Make sure registering works"""
    r = register('usserer11', 'default')
    assert r.ok
    r = register('usserer11', 'default')
    assert 'The username is already taken' in r.text
    r = register('', 'default')
    assert 'You have to enter a username' in r.text
    r = register('meh', '')
    assert 'You have to enter a password' in r.text
    r = register('meh', 'foo', email='broken')
    assert 'You have to enter a valid email address' in r.text

# Working
def test_login_logout():
    """Make sure logging in and logging out works"""
    r, http_session = register_and_login('user1', 'default')
    assert r.ok
    r = logout(http_session)
    assert r.ok
    r, _ = login('user1', 'wrongpassword')
    assert 'Invalid password' in r.text
    r, _ = login('user2', 'wrongpassword')
    assert 'Invalid username' in r.text

# Working
def test_message_recording():
    """Check if adding messages works"""
    r, http_session = register_and_login('foo', 'default')
    add_message(http_session, 'test message 1')
    add_message(http_session, '<test message 2>')
    r = http_session.get(f'{BASE_URL}/timeline')
    data = json.dumps(r.json())
    assert 'test message 1' in data
    assert '<test message 2>' in data

# Working
def test_timelines():
    """Make sure that timelines work"""
    _, http_session = register_and_login('foo', 'default')
    add_message(http_session, 'the message by foo')
    logout(http_session)
    _, http_session = register_and_login('bar', 'default')
    add_message(http_session, 'the message by bar')
    r = http_session.get(f'{BASE_URL}/public')
    data = json.dumps(r.json())
    assert 'the message by foo' in data
    assert 'the message by bar' in data

    # bar's timeline should just show bar's message
    r = http_session.get(f'{BASE_URL}/timeline')
    data = json.dumps(r.json())
    assert 'the message by foo' not in data
    assert 'the message by bar' in data

    # now let's follow foo
    r = http_session.post(f'{BASE_URL}/foo/follow', allow_redirects=True)
    assert r.ok

    # we should now see foo's message
    r = http_session.get(f'{BASE_URL}/timeline')
    data = json.dumps(r.json())
    assert 'the message by foo' in data
    assert 'the message by bar' in data

    # but on the user's page we only want the user's message
    r = http_session.get(f'{BASE_URL}/timeline/bar')
    data = json.dumps(r.json())
    assert 'the message by foo' not in data
    assert 'the message by bar' in data
    r = http_session.get(f'{BASE_URL}/timeline/foo')
    data = json.dumps(r.json())
    assert 'the message by foo' in data
    assert 'the message by bar' not in data

    # now unfollow and check if that worked
    r = http_session.post(f'{BASE_URL}/foo/unfollow', allow_redirects=True)
    assert r.ok
    r = http_session.get(f'{BASE_URL}/timeline')
    data = json.dumps(r.json())
    assert 'the message by foo' not in data
    assert 'the message by bar' in data

test_register()
test_login_logout()
test_message_recording()
test_timelines()
