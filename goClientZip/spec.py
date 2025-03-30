
import json

def get_refs(json_obj):
    result = set()
    if isinstance(json_obj, dict):
        for key, value in json_obj.items():
            if key == "$ref":
                result.add(value.split('/')[-1])
            else:
                result = result | get_refs(value)
    elif isinstance(json_obj, list):
        for item in json_obj:
            result = result | get_refs(item)
    return result

def filter_by_paths(json_obj, paths):
    ret = {}
    for key in json_obj:
        if key in paths:
            ret[key] = json_obj[key]
            for op in ret[key]:
                if op not in ['get', 'put', 'post', 'delete', 'options', 'head', 'patch']:
                    continue
                # print(ret[key][op])
                ret[key][op]['operationId'] = addOpId(key, op)
    return ret

def getnextqueue(json_obj_defs, refs):
    next_queue = set()
    for ref in refs:
        next_queue = next_queue | get_refs(json_obj_defs[ref])
    return next_queue

def getAllRefsRec(json_obj_defs, refs):
    print("The initial queue is: ", refs)
    it = 0
    ret = set()
    while(len(refs)>0):
        print("iteration: ", str(it), " =============== queue: ", refs)
        ret = ret | refs
        nextlevel = getnextqueue(json_obj_defs, refs)
        refs = nextlevel - ret
        it+=1
    return ret

def filter_models_by_paths(json_obj_defs, top_level_refs):
    
    x = getAllRefsRec(json_obj_defs, top_level_refs)
    print("All refs are: ", x)

    redundant_keys = set(json_obj_defs.keys()) - x
    for key in redundant_keys:
        del json_obj_defs[key]
    return json_obj_defs

def addOpId(path, opType):
    components = path.split('/')
    if len(components) == 3:
        op = opType + '_' + components[1] + '_by_id'
    else :
        op = opType + '_all_' + components[1] + 's'
    return op

def get_openapi(file_path, paths):
    with open(file_path, 'r') as file:
        json_obj = json.load(file)
    # json_obj['paths'] = filter_by_paths(json_obj['paths'], ['/volume', "/volume/{id}"])
    json_obj['paths'] = filter_by_paths(json_obj['paths'], paths)
    
    top_level_refs = get_refs(json_obj['paths'])
    print("The top level models are: ", top_level_refs)

    filter_models_by_paths(json_obj['definitions'], top_level_refs)
    print("The number of models is: ", len(json_obj['definitions'].keys()))

    return json_obj

# Example usage
filtered_json = get_openapi('/root/terraform-provider-powerstore/goClientZip/spec_4_1.json', ['/volume_group', "/volume_group/{id}"])
# write to file
with open('/root/terraform-provider-powerstore/goClientZip/spec_4_1_filtered.json', 'w') as outfile:
    json.dump(filtered_json, outfile)
