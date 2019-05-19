import React from 'react';
import { shallow } from 'enzyme';
import Label from '../Label';

it('renders without crashing', () => {
  shallow(<Label />);
});
