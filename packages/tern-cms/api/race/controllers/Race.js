'use strict';

/**
 * Race.js controller
 *
 * @description: A set of functions called "actions" for managing `Race`.
 */

module.exports = {

  /**
   * Retrieve race records.
   *
   * @return {Object|Array}
   */

  find: async (ctx) => {
    if (ctx.query._q) {
      return strapi.services.race.search(ctx.query);
    } else {
      return strapi.services.race.fetchAll(ctx.query);
    }
  },

  /**
   * Retrieve a race record.
   *
   * @return {Object}
   */

  findOne: async (ctx) => {
    return strapi.services.race.fetch(ctx.params);
  },

  /**
   * Count race records.
   *
   * @return {Number}
   */

  count: async (ctx) => {
    return strapi.services.race.count(ctx.query);
  },

  /**
   * Create a/an race record.
   *
   * @return {Object}
   */

  create: async (ctx) => {
    return strapi.services.race.add(ctx.request.body);
  },

  /**
   * Update a/an race record.
   *
   * @return {Object}
   */

  update: async (ctx, next) => {
    return strapi.services.race.edit(ctx.params, ctx.request.body) ;
  },

  /**
   * Destroy a/an race record.
   *
   * @return {Object}
   */

  destroy: async (ctx, next) => {
    return strapi.services.race.remove(ctx.params);
  },

  /**
   * Add relation to a/an race record.
   *
   * @return {Object}
   */

  createRelation: async (ctx, next) => {
    return strapi.services.race.addRelation(ctx.params, ctx.request.body);
  },

  /**
   * Update relation to a/an race record.
   *
   * @return {Object}
   */

  updateRelation: async (ctx, next) => {
    return strapi.services.race.editRelation(ctx.params, ctx.request.body);
  },

  /**
   * Destroy relation to a/an race record.
   *
   * @return {Object}
   */

  destroyRelation: async (ctx, next) => {
    return strapi.services.race.removeRelation(ctx.params, ctx.request.body);
  }
};
