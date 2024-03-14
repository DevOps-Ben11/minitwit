import React, { ReactNode } from 'react';
import { Link } from 'react-router-dom';

/*
    This component is used to show the top bar and bottom bar of minitwit to the user.
    The children is a component like the register or login page, it will be placed between the top and bottom bar.

    Things that doesn't work right now:
        - The user is alway null because I haven't found how to login someone.

    This component works for now and only thing missing is the implementation of a connected user,
    but that is something that should not be handle here.
*/

const Layout: React.FC<{ children: ReactNode, user: { username: string } | null }> = ({ user, children }) => {
  return (
    <div className="page">
      <h1>MiniTwit</h1>
      <div className="navigation">
        {user ? (
          <>
            <Link to="/timeline">my timeline</Link> |
            <Link to="/public_timeline">public timeline</Link> |
            <Link to="/logout">sign out [{user.username}]</Link>
          </>
        ) : (
          <>
            <Link to="/public_timeline">public timeline</Link> |
            <Link to="/register">sign up</Link> |
            <Link to="/login">sign in</Link>
          </>
        )}
      </div>
      <div className="body">
        {children /* The children "other component" is here! */ }
      </div>
      <div className="footer">
        MiniTwit &mdash; A Go Application
      </div>
    </div>
  );
};

export default Layout;
