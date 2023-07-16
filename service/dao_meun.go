package app

import (
	"context"
	"github.com/ml444/gkit/dbx"
	"github.com/ml444/gkit/listoption"
	log "github.com/ml444/glog"
	"gorm.io/gorm"

	"github.com/JIAWea/erpServer/api/erp"
	"github.com/JIAWea/erpServer/internal/db"
)

func init() {
	db.RegisterModel(&erp.ModelUser{}, &erp.ModelRole{}, &erp.ModelMenu{}, &erp.ModelUserRole{}, &erp.ModelRoleMenu{})
}

var dbMenu = NewTMenu(db.Db())

type TMenu struct {
	db    *gorm.DB
	model *erp.ModelMenu
}

func NewTMenu(db *gorm.DB) *TMenu {
	return &TMenu{
		db:    db,
		model: &erp.ModelMenu{},
	}
}

func (d *TMenu) newScope() *dbx.Scope {
	if d.db == nil {
		d.db = db.Db()
	}
	return dbx.NewScope(d.db, &erp.ModelMenu{})
}

func (d *TMenu) Create(ctx context.Context, m *erp.ModelMenu) error {
	return d.newScope().Create(ctx, &m)
}

func (d *TMenu) Update(ctx context.Context, m *erp.ModelMenu, whereMap map[string]interface{}) error {
	return d.newScope().Where(whereMap).Update(&m)
}

func (d *TMenu) DeleteById(ctx context.Context, pk uint64) error {
	return d.newScope().Delete(&erp.ModelMenu{}, pk)
}

func (d *TMenu) DeleteByWhere(ctx context.Context, whereMap map[string]interface{}) error {
	return d.newScope().Delete(&erp.ModelMenu{}, whereMap)
}

func (d *TMenu) GetOne(ctx context.Context, pk uint64) (*erp.ModelMenu, error) {
	var m erp.ModelMenu
	err := d.newScope().SetNotFoundErr(erp.ErrNotFoundMenu).First(&m, pk)
	return &m, err
}

func (d *TMenu) ListWithListOption(ctx context.Context, listOption *listoption.ListOption, whereOpts interface{}) ([]*erp.ModelMenu, *listoption.Paginate, error) {
	var err error
	scope := d.newScope().Where(whereOpts)
	if listOption != nil {

		err = listoption.NewProcessor(listOption).
			AddUint32(erp.ListMenuReq_ListOptType, func(val uint32) error {
				scope.Eq(dbType, val)
				return nil
			}).
			Process()
		if err != nil {
			log.Error(err.Error())
			return nil, nil, err
		}

	}

	var menuList []*erp.ModelMenu
	var paginate *listoption.Paginate
	paginate, err = scope.PaginateQuery(listOption, &menuList)
	if err != nil {
		log.Errorf("err: %v", err)
		return nil, nil, err
	}

	return menuList, paginate, nil
}

func (d *TMenu) getMenuTree(parentId uint64, menus []*erp.ListMenuTreeRsp_Node) []*erp.ListMenuTreeRsp_Node {
	tree := make([]*erp.ListMenuTreeRsp_Node, 0)
	for _, m := range menus {
		if m.ParentId == parentId {
			children := d.getMenuTree(m.Id, menus)
			m.Children = children
			tree = append(tree, m)
		}
	}
	return tree
}

func (d *TMenu) GetMenuTree(ctx context.Context, listOption *listoption.ListOption) ([]*erp.ListMenuTreeRsp_Node, error) {
	var (
		err   error
		nodes []*erp.ListMenuTreeRsp_Node
	)
	scope := d.newScope().Eq(dbStatus, uint32(erp.ModelUser_StatusEnable))
	if listOption != nil {
		listOption.SkipCount = true
		err = listoption.NewProcessor(listOption).
			AddUint32(erp.ListMenuTreeReq_ListOptType, func(val uint32) error {
				scope.Eq(dbType, val)
				return nil
			}).
			AddString(erp.ListMenuTreeReq_ListOptComponent, func(val string) error {
				scope.Eq(dbComponent, val)
				return nil
			}).
			Process()
		if err != nil {
			log.Error(err.Error())
			return nil, err
		}
	}

	var menuList []*erp.ModelMenu
	_, err = scope.PaginateQuery(listOption, &menuList)
	if err != nil {
		log.Errorf("err: %v", err)
		return nil, err
	}

	if len(menuList) == 0 {
		return nodes, nil
	}

	for _, v := range menuList {
		nodes = append(nodes, &erp.ListMenuTreeRsp_Node{
			Id:       v.Id,
			Name:     v.Name,
			ParentId: v.ParentId,
		})
	}

	tree := d.getMenuTree(0, nodes)

	return tree, nil
}

func (d *TRoleMenu) GetRoleMenuIdList(ctx context.Context, typ uint32, roleId uint64) ([]uint64, error) {
	var idList []uint64

	var roleMenuList []*erp.ModelRoleMenu
	err := d.newScope().
		Joins("INNER JOIN erp_menu m ON m.id = erp_role_menu.menu_id").
		Where("erp_role_menu.role_id = ?", roleId).
		Where("m.type = ?", typ).
		Where("m.status = ?", uint32(erp.ModelMenu_StatusEnable)).
		Find(&roleMenuList)
	if err != nil {
		return nil, err
	}

	for _, menu := range roleMenuList {
		idList = append(idList, menu.MenuId)
	}

	return idList, nil
}

func (d *TMenu) getUserMenuTree(parentId uint64, menus []*erp.MenuTree) []*erp.MenuTree {
	tree := make([]*erp.MenuTree, 0)
	for _, m := range menus {
		if m.ParentId == parentId {
			children := d.getUserMenuTree(m.Id, menus)
			m.Children = children
			tree = append(tree, m)
		}
	}
	return tree
}

func (d *TMenu) GetUserMenuList(ctx context.Context, userId uint64) ([]*erp.MenuTree, error) {
	var menuList []*erp.ModelMenu
	err := d.newScope().
		Joins("INNER JOIN erp_role_menu r ON r.menu_id = erp_menu.id").
		Where("erp_menu.type = ?", uint32(erp.ModelMenu_TypeMenu)).
		Where("erp_menu.status = ?", uint32(erp.ModelMenu_StatusEnable)).
		Where("r.role_id IN (SELECT role_id FROM erp_user_role WHERE user_id = ?)", userId).
		Find(&menuList)
	if err != nil {
		return nil, err
	}

	var data []*erp.MenuTree
	for _, v := range menuList {
		data = append(data, &erp.MenuTree{
			Id:        v.Id,
			Name:      v.Name,
			Icon:      v.Icon,
			Path:      v.Path,
			Redirect:  v.Redirect,
			Component: v.Component,
			IsHidden:  v.IsHidden,
			ParentId:  v.ParentId,
		})
	}
	if len(data) == 0 {
		return data, nil
	}

	return d.getUserMenuTree(0, data), nil
}

func (d *TMenu) GetAllUserMenuList(ctx context.Context) ([]*erp.MenuTree, error) {
	var menuList []*erp.ModelMenu
	err := d.newScope().
		Where("erp_menu.type = ?", uint32(erp.ModelMenu_TypeMenu)).
		Where("erp_menu.status = ?", uint32(erp.ModelMenu_StatusEnable)).
		Find(&menuList)
	if err != nil {
		return nil, err
	}

	var data []*erp.MenuTree
	for _, v := range menuList {
		data = append(data, &erp.MenuTree{
			Id:        v.Id,
			Name:      v.Name,
			Icon:      v.Icon,
			Path:      v.Path,
			Redirect:  v.Redirect,
			Component: v.Component,
			IsHidden:  v.IsHidden,
			ParentId:  v.ParentId,
		})
	}
	if len(data) == 0 {
		return data, nil
	}

	return d.getUserMenuTree(0, data), nil
}
