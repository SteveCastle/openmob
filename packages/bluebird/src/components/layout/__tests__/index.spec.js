import React from 'react';
import { shallow } from 'enzyme';
import index from '../index';

it('renders without crashing', () => {
  shallow(<index />);
});
