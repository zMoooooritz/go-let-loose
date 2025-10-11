---@meta

-- Plugin Lifecycle Interface
-- These functions define the structure that plugins should implement

---Plugin initialization function
---Called once when the plugin is first loaded, before Run() is called.
---Use this for setting up initial state, printing startup messages, or other one-time setup.
---@return nil
function Init() end

---Main plugin execution function
---Called after Init() completes successfully. This is where the main plugin logic should reside.
---Set up event handlers, register callbacks, and implement the core functionality here.
---This function should not block - use event handlers for ongoing functionality.
---@return nil
function Run() end

---Plugin cleanup function
---Called when the plugin is being stopped or the system is shutting down.
---Use this for cleanup operations, saving state, or other shutdown tasks.
---MUST call exit() at the end of this function to properly terminate the plugin.
---@return nil
function Stop() end

---Exit the plugin gracefully
---This function must be called at the end of the Stop() function to properly terminate the plugin.
---Signals to the plugin manager that the plugin has finished its cleanup and can be safely unloaded.
---@return nil
function exit() end