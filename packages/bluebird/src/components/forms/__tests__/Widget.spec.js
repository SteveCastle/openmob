import React from 'react';
import { shallow } from 'enzyme';
import Widget from '../Widget';

it('renders without crashing', () => {
  shallow(<Widget />);
});
