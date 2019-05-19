import React from 'react';
import { shallow } from 'enzyme';
import SimpleBlogPost from '../SimpleBlogPost';

it('renders without crashing', () => {
  shallow(<SimpleBlogPost />);
});
