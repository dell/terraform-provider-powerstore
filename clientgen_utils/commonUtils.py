
import json

def _get_refs(json_obj: dict) -> set:
    """
    Gets all models referenced in an openapi spec section.
    The section could be a model or a path.

    Args:
        json_obj (dict): The JSON object.
    
    Returns:
        set: The set of referenced models
    """
    result = set()
    if isinstance(json_obj, dict):
        for key, value in json_obj.items():
            if key == "$ref":
                result.add(value.split('/')[-1])
            else:
                result = result | _get_refs(value)
    elif isinstance(json_obj, list):
        for item in json_obj:
            result = result | _get_refs(item)
    return result

def _filter_by_paths(json_obj, paths):
    ret = {}
    for key in json_obj:
        if key in paths:
            ret[key] = json_obj[key]
    return ret

def _getnextqueue(json_obj_defs, refs):
    next_queue = set()
    for ref in refs:
        next_queue = next_queue | _get_refs(json_obj_defs[ref])
    return next_queue

def _getAllRefsRec(all_models, refs):
    print("The initial queue is: ", refs)
    it = 0
    ret = set()
    while(len(refs)>0):
        print("iteration: ", str(it), " =============== queue: ", refs)
        ret = ret | refs
        nextlevel = _getnextqueue(all_models, refs)
        refs = nextlevel - ret
        it+=1
    return ret

def _get_all_required_models(all_models: dict, top_level_models: set) -> dict:
    """
    This function filters the JSON object definitions to only include the top-level references and their dependencies.

    Args:
        all_models (dict): All model definitions.
        top_level_refs (set): The top-level models.

    Returns:
        dict: The filtered JSON object definitions.
    """
    
    # Get all references recursively
    all_refs = _getAllRefsRec(all_models, top_level_models)
    print("All refs are: ", all_refs)

    # Find the redundant keys by taking the difference between all keys and the references
    redundant_keys = set(all_models.keys()) - all_refs
    
    # Delete the redundant keys from the JSON object definitions
    for key in redundant_keys:
        del all_models[key]
    
    return all_models

def ProcessOpenapiSpec(file_path, paths):
    with open(file_path, 'r') as file:
        json_obj = json.load(file)
    json_obj['paths'] = _filter_by_paths(json_obj['paths'], paths)
    
    top_level_refs = _get_refs(json_obj['paths'])
    print("The top level models are: ", top_level_refs)

    _get_all_required_models(json_obj['definitions'], top_level_refs)
    print("The number of models is: ", len(json_obj['definitions'].keys()))

    return json_obj
