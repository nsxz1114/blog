"use client";
import { useEffect, useState } from "react";
import { ProductItem, Menu, MenuItem } from "../../ui/navbar-menu";
import { navList, navListType } from "@/api/system";

export function NavbarDemo() {
  const [navItems, setNavItems] = useState<navListType[]>([]);
  useEffect(() => {
    const fetchNavList = async () => {
      const res = await navList();
      setNavItems(res.data.list);
    };
    fetchNavList();
  }, []);
  const [active, setActive] = useState<string | null>(null);
  return (
    <div>
      <Menu setActive={setActive}>
        {navItems.map((navItem) => (
          <MenuItem
            key={navItem.item}
            setActive={setActive}
            active={active}
            item={navItem.item}
          >
            <ProductItem
              title={navItem.title}
              href={navItem.href}
              src={navItem.src}
              description={navItem.description}
            />
          </MenuItem>
        ))}
      </Menu>
    </div>
  );
}
