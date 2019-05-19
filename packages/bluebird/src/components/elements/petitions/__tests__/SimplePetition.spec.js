import React from 'react';
import { shallow } from 'enzyme';
import SimplePetition from '../SimplePetition';

it('renders without crashing', () => {
  shallow(<SimplePetition />);
});
