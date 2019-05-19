import React from 'react';
import { shallow } from 'enzyme';
import Column from '../Column';

it('renders without crashing', () => {
  shallow(<Column />);
});
