import React from 'react';
import { shallow } from 'enzyme';
import SimpleFooter from '../SimpleFooter';

it('renders without crashing', () => {
  shallow(<SimpleFooter />);
});
