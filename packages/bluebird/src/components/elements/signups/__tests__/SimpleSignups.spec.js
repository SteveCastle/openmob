import React from 'react';
import { shallow } from 'enzyme';
import SimpleSignups from '../SimpleSignups';

it('renders without crashing', () => {
  shallow(<SimpleSignups />);
});
