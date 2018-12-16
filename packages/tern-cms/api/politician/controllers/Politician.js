'use strict';

/**
 * Politician.js controller
 *
 * @description: A set of functions called "actions" for managing `Politician`.
 */

module.exports = {

  /**
   * Retrieve politician records.
   *
   * @return {Object|Array}
   */

  find: async (ctx) => {
    if (ctx.query._q) {
      return strapi.services.politician.search(ctx.query);
    } else {
      return strapi.services.politician.fetchAll(ctx.query);
    }
  },

  /**
   * Retrieve a politician record.
   *
   * @return {Object}
   */

  findOne: async (ctx) => {
    return strapi.services.politician.fetch(ctx.params);
  },

  /**
   * Count politician records.
   *
   * @return {Number}
   */

  count: async (ctx) => {
    return strapi.services.politician.count(ctx.query);
  },

  /**
   * Create a/an politician record.
   *
   * @return {Object}
   */

  create: async (ctx) => {
    return strapi.services.politician.add(ctx.request.body);
  },

  /**
   * Update a/an politician record.
   *
   * @return {Object}
   */

  update: async (ctx, next) => {
    return strapi.services.politician.edit(ctx.params, ctx.request.body) ;
  },

  /**
   * Destroy a/an politician record.
   *
   * @return {Object}
   */

  destroy: async (ctx, next) => {
    return strapi.services.politician.remove(ctx.params);
  },

  /**
   * Add relation to a/an politician record.
   *
   * @return {Object}
   */

  createRelation: async (ctx, next) => {
    return strapi.services.politician.addRelation(ctx.params, ctx.request.body);
  },

  /**
   * Update relation to a/an politician record.
   *
   * @return {Object}
   */

  updateRelation: async (ctx, next) => {
    return strapi.services.politician.editRelation(ctx.params, ctx.request.body);
  },

  /**
   * Destroy relation to a/an politician record.
   *
   * @return {Object}
   */

  destroyRelation: async (ctx, next) => {
    return strapi.services.politician.removeRelation(ctx.params, ctx.request.body);
  }
};
