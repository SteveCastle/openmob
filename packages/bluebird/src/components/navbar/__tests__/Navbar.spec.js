import React from 'react';
import { shallow } from 'enzyme';
import Navbar from '../Navbar';

it('renders without crashing', () => {
  shallow(<Navbar />);
});
