import React from 'react';
import { shallow } from 'enzyme';
import Container from '../Container';

it('renders without crashing', () => {
  shallow(<Container />);
});
