import React from 'react';
import { shallow } from 'enzyme';
import SimpleBoycott from '../SimpleBoycott';

it('renders without crashing', () => {
  shallow(<SimpleBoycott />);
});
